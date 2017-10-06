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


n.n.n / 2015-07-18
==================

  * Fix: Dockerfile had wrong argument usage

1.3.0 / 2015-05-29
==================

  * Replaced codegangsta/cli by spf13/cobra
  * Fixed some wording / URLs in README

1.2.2 / 2015-05-10
==================

  * Fix: Blacklisted characters were not excluded from passwords

1.2.1 / 2015-05-10
==================

  * Added links to extension & documentation
  * Fix: Load bootstrap theme via HTTPs

1.2.0 / 2015-05-10
==================

  * Included web frontend into server

1.1.0 / 2015-05-03
==================

  * Added test for password generation
  * Moved error handling for short passwords to library

1.0.2 / 2015-05-03
==================

  * Fix: Ensured minimal length of passwords
  * Switched to https protocol
  * Added additional insecure password test
  * Added API documentation to README

1.0.1 / 2015-05-02
==================

  * Dockerized for API deployment
  * Added travis build and README

1.0.0 / 2015-05-02
==================

  * Initial running version