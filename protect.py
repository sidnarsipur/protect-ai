from docopt import docopt
from logging import error

import sys
import csv

doc = """Usage: protect.py SITEMAP [--allowSEO] [--blockAll] [--blockTraining] [--allowCrawler=Crawlers] [--disallowCrawler=Crawlers] [--allowFile=FileTypes] [--disallowFile=FileTypes] [--disallowDir=FileTypes]
"""

blockAll = ['User-Agent: * \n', 'Disallow: / \n']

def validate_sitemap(sitemap: str):
  if not sitemap.endswith(".xml"):
    raise Exception("ERROR: Sitemap must end with .xml")
  elif not "www." in sitemap in sitemap:
    raise Exception("ERROR: Sitemap must be in the form www.[some string].com")

  return sitemap

def parse_args(allowed_crawlers, disallowed_crawlers, allowed_files, disallowed_files, disallowed_dirs):
  a_crawlers = allowed_crawlers.split(',') if allowed_crawlers else []
  d_crawlers = disallowed_crawlers.split(',') if disallowed_crawlers else []
  a_files = allowed_files.split(',') if allowed_files else []
  d_files = disallowed_files.split(',') if disallowed_files else []
  d_dirs = disallowed_dirs.split(',') if disallowed_dirs else []

  return a_crawlers, d_crawlers, a_files, d_files, d_dirs

def read_crawlers_csv(filename="Crawlers.csv"):
  crawlers = []
  with open(filename, "r") as csvfile:
    reader = csv.DictReader(csvfile)
    for row in reader:
      crawlers.append({"User-Agent": row["User-Agent"], "Type": row["Type"]})
  return crawlers

def get_seo_crawlers(crawlers):
  seo_crawler = []

  for crawler in crawlers:
    if crawler["Type"] == "seo":
      seo_crawler.append(crawler["User-Agent"])

def get_training_crawlers(crawlers):
  training_crawler = []

  for crawler in crawlers:
    if crawler["Type"] == "training":
      training_crawler.append(crawler["User-Agent"])
  
  return training_crawler

def get_crawler_string(crawler='', d_dirs=[], a_files=[], d_files=[], type=''):

  if crawler != '':
    user_agent_string = '\nUser-Agent: '  + crawler + '\n'
    lines = [user_agent_string]

  if type == 'seo' or type == 'allow':
    lines.append('Allow: /' + '\n')

  elif not d_dirs:
    lines.append('Disallow: /' + '\n')

  for dirc in d_dirs:
    lines.append('Disallow: /' + dirc + '\n')

  for file in a_files:
    lines.append('Allow: /*.' + file + '$' + '\n')
  
  for file in d_files:
    lines.append('Disallow: /*.' + file + '$' + '\n')

  return lines

def write_robot(args, crawlers):
  sitemap = validate_sitemap(args['SITEMAP'])
  a_crawlers, d_crawlers, a_files, d_files, d_dirs = parse_args(args['--allowCrawler'],args['--disallowCrawler'], args['--allowFile'], args['--disallowFile'], args['--disallowDir'])
  
  seo_crawlers = get_seo_crawlers(crawlers)
  training_crawlers = get_training_crawlers(crawlers)

  with open("robots.txt", "w+") as robot:
    if args['--blockAll']:
      robot.writelines(blockAll)
    else:
      robot.writelines(get_crawler_string("*", d_dirs, [], d_files, 'allow'))

    for crawler in crawlers:
      if args['--allowSEO']:
        if crawler['User-Agent'] in seo_crawlers:
          robot.writelines(get_crawler_string(crawler['User-Agent'], d_dirs, a_files, d_files, 'seo'))
      if args['--blockTraining']:
        if crawler['User-Agent'] in training_crawlers:
          robot.writelines(get_crawler_string(crawler['User-Agent']))

    for crawler in a_crawlers:
      robot.writelines(get_crawler_string(crawler, d_dirs, a_files, d_files, 'allow'))
    
    for crawler in d_crawlers:
      robot.writelines(get_crawler_string(crawler))
    
    robot.writelines('\nSitemap: ' + sitemap)

if __name__ == '__main__':
  args = docopt(doc, help=False)
  crawlers = read_crawlers_csv()

  write_robot(args, crawlers)




