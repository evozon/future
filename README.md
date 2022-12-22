# Future

Future is a library that facilitates the upgrade process of PHP projects. It can help you in 3 major areas:

### 1. Platform Upgrade

Future will test if your project is able to run using the latest platform configuration (e.g. the latest PHP version) and will provide a detailed summary of any encountered issue.

### 2. Dependencies Upgrade

On top of the platform upgrade, it will test if your project is able to run using the latest version of your direct Composer dependencies and will provide a list of blockers that stop you from upgrading.

### 3. Code Upgrade

Finally, Future will test if your project is compatible with the latest coding standards with the help of [Rector](https://github.com/rectorphp/rector).

## Setup

### Prerequisites
* A Composer-based Project
* A Continuous Integration Pipeline
* A Test Suite

### Install
```bash
composer require --dev evozon/future
```

### Configure

Please refer to the links below for information on how to configure Future:

* [Configuring Future for GitLab](https://github.com/evozon/future/blob/master/docs/GITLAB.md)
* [Configuring Future for GitHub](https://github.com/evozon/future/blob/master/docs/GITHUB.md)

## Running Future

After you have installed and configured Future, run the pipeline and check the output of the `future-proofing` job.

## Contribute
Please refer to [CONTRIBUTING.md](https://github.com/evozon/future/blob/master/CONTRIBUTING.md) for information on how to contribute to Future.

## Recommendations
Future can be used to test if you can upgrade everything at once: the latest PHP version, the latest versions for your Composer dependencies, and the latest codebase standards. We advise against this since it may lead to massive amounts of changes that are hard to review and test.

We recommend splitting the upgrade process into stages. Start with the PHP version upgrade (one minor version at a time), continue with the dependencies upgrade and finish with the codebase changes. This way you can have smaller PRs that are easier to review and test.

<hr/>

Thanks to [Rector](https://github.com/rectorphp/rector) for providing the tool that does the heavy lifting. Future uses Rector to upgrade the codebase to the latest coding standards.
