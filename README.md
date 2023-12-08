About
==========

Protect AI blocks unauthorized web crawlers from accessing pages and files on your website by generating a custom [robots.txt](https://developers.google.com/search/docs/crawling-indexing/robots/create-robots-txt) file. 

Installation
==============

Package Coming soon!

For now - clone (or download) the repo and see the commands below.

Usage
===========

`protect.py SITEMAP [--allowSEO] [--blockAll] [--blockTraining] [--allowCrawler=Crawlers] [--disallowCrawler=Crawlers] [--allowFile=FileTypes] [--disallowFile=FileTypes] [--disallowDir=FileTypes]`

**Options**:

- `SITEMAP`: Path to your website's sitemap file.
- `--allowSEO`: (Optional, Default=false) Set to true to allow search engine (Google, Bing) crawlers access.
- `--blockAll`: (Optional, Default=false) Block all crawlers except those explicitly allowed.
- `--blockTraining`: (Optional, Default=false) Block crawlers known to use data for model training (OpenAI etc.).
- `--allowCrawler`: (Optional) List of crawler names to allow access.
- `--disallowCrawler`: (Optional) List of crawler names to deny access. Recommended if `blockAll` is false.
- `--allowFile`: (Optional) List of file extensions to allow access.
- `--disallowFile`: (Optional) List of file extensions to deny access to all crawlers.
- `--disallowDir`: (Optional) List of directory paths (from the website root) to deny access to all crawlers.

Example
===========

- Allow all crawlers access to all files and directories:

`python3 Protect.py www.hello.xml`

- Block all crawlers but allow Googlebot:

`protect.py my_sitemap.xml --blockAll --allowCrawler=Googlebot`

- Block all crawlers but allow msnbot, DiscordBot and seo crawlers:

`protect.py my_sitemap.xml --allowSEO --blockAll --allowCrawler=msnbot,DiscordBot`

- Block all crawlers but allow msnbot, DiscordBot and seo crawlers but block the pics/ directory:

`protect.py my_sitemap.xml --allowSEO --blockAll --allowCrawler=msnbot,DiscordBot --disallowDir=pics/`

- Allow all crawlers access to all files except PDF and JPG files:

`python3 Protect.py www.hello.xml --disallowFile=pdf,jpg`

- Allow only SEO crawlers

`python3 Protect.py www.hello.xml --blockAll --allowSEO`

- Block only training crawlers

`python3 Protect.py www.hello.xml --blockTraining`

