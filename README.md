<p align="center">
    <img alt="Pushover Actions Logo" src="https://raw.githubusercontent.com/Clivern/pushover-actions/master/assets/images/logo.png" height="130" />
    <h3 align="center">Pushover Actions</h3>
    <p align="center">Pushover A Notifications for Github Repository Changes</p>
    <p align="center">
        <a href="https://travis-ci.com/Clivern/pushover-actions"><img src="https://travis-ci.com/Clivern/pushover-actions.svg?branch=master"></a>
        <a href="https://github.com/Clivern/pushover-actions/releases"><img src="https://img.shields.io/badge/Version-0.0.1-red.svg"></a>
        <a href="https://github.com/Clivern/pushover-actions/blob/master/LICENSE"><img src="https://img.shields.io/badge/LICENSE-Apache--2.0-orange.svg"></a>
    </p>
</p>


## Documentation

#### Pushover as a Github Action:

To install pushover github action, add the following to your `workflow.yml`.

```yml
name: workflow_name

on:
    # may vary based on modules enabled

jobs:
  pushover-actions:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2 # required to clone your code
      - name: pushover-actions
        uses: clivern/pushover-actions@master
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          PUSHOVER_TOKEN: ${{ secrets.PUSHOVER_TOKEN }}
          PUSHOVER_USER: ${{ secrets.PUSHOVER_USER }}
```


## Versioning

For transparency into our release cycle and in striving to maintain backward compatibility, Pushover Actions is maintained under the [Semantic Versioning guidelines](https://semver.org/) and release process is predictable and business-friendly.

See the [Releases section of our GitHub project](https://github.com/clivern/pushover-actions/releases) for changelogs for each release version of Pushover Actions. It contains summaries of the most noteworthy changes made in each release.


## Bug tracker

If you have any suggestions, bug reports, or annoyances please report them to our issue tracker at https://github.com/clivern/pushover-actions/issues


## Security Issues

If you discover a security vulnerability within Pushover Actions, please send an email to [hello@clivern.com](mailto:hello@clivern.com)


## Contributing

We are an open source, community-driven project so please feel free to join us. see the [contributing guidelines](CONTRIBUTING.md) for more details.


## License

© 2020, Clivern. Released under [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0).

**Pushover-Actions** is authored and maintained by [@Clivern](https://github.com/clivern).
