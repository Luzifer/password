# 2.0.0 / 2019-02-28

  - Fix broken vendoring for library users (closes #4)
    - **Breaking:** Move command to `cmd/password` subdir
  - **Breaking:** Drop support for Alfred workflow
  - Add support for Go 1.11+ modules

# 1.13.1 / 2019-02-01

  * Fix: Add missing dependencies

# 1.13.0 / 2019-01-31

  * Add ability to check the passwords against HIBP database

# 1.12.1 / 2019-01-01

  * Update README
  * Build on more Go versions

# 1.12.0 / 2019-01-01

  * Port application to Bootstrap 4
  * Fix: Copy&Paste error

# 1.11.1 / 2018-12-30

  * Fix: Load JS earlier

# 1.11.0 / 2018-12-30

  * Port from coffeescript to plain ES6
  * Ensure fonts are loaded from local instead of gfonts
  * Enable gzip compression on assets
  * Fix: RawGit is closing down, use Github URL
  * Update vendored libraries

# 1.10.3 / 2018-12-30

  * Fix: Mime guessing takes only the extension

# 1.10.2 / 2018-12-30

  * Update Dockerfile
  * Fix: Serve needs to provide proper mime types

# 1.10.1 / 2018-03-21

  * Fix: Workflow building broken by zipping artifacts
  * Beautify History.md
  * Add hint for XKCD style passwords

# 1.10.0 / 2018-03-21

  * Make generated passwords actionable again (#3 - Thanks to @funkjedi)
  * Update workflow library

# 1.9.0 / 2017-10-31

  * Add commandline flag to generate multiple passwords

# 1.8.0 / 2017-10-31

  * Add XKCD style password generator in library
  * Add ability to generate XKCD style passwords to tool
  * Add ability for XKCD style passwords to workflow

# 1.7.1 / 2017-10-06

  * Fix: Vendor new dependencies

# 1.7.0 / 2017-10-06

  * Add option to workflow to copy hashed versions
  * Implement JSON output with password hashes

# 1.6.2 / 2017-09-22

  * Fix location of workflow in the frontend
  * Add instructions for Alfred workflow
  * Move workflow building to extra script

# 1.6.1 / 2017-09-22

  * Fix: Need to download submodules before packaging

# 1.6.0 / 2017-09-22

  * Rebuilt workflow for better design and auto-update

# 1.5.0 / 2017-09-22

  * Include workflow building into build process
  * Fix: Travis build broken through different tools

# 1.4.0 / 2017-09-22

  * Add buttons to README
  * Replace Apache license stub with proper license
  * Remove references to GoBuilder
  * Add Github release publishing
  * Remove static versioning
  * Vendor dependencies
  * Update Dockerfile to golang:alpine template
  * Fix: Dockerfile had wrong argument usage

# 1.3.0 / 2015-05-29

  * Replaced codegangsta/cli by spf13/cobra
  * Fixed some wording / URLs in README

# 1.2.2 / 2015-05-10

  * Fix: Blacklisted characters were not excluded from passwords

# 1.2.1 / 2015-05-10

  * Added links to extension & documentation
  * Fix: Load bootstrap theme via HTTPs

# 1.2.0 / 2015-05-10

  * Included web frontend into server

# 1.1.0 / 2015-05-03

  * Added test for password generation
  * Moved error handling for short passwords to library

# 1.0.2 / 2015-05-03

  * Fix: Ensured minimal length of passwords
  * Switched to https protocol
  * Added additional insecure password test
  * Added API documentation to README

# 1.0.1 / 2015-05-02

  * Dockerized for API deployment
  * Added travis build and README

# 1.0.0 / 2015-05-02

  * Initial running version
