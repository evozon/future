# Future

Future is a library that facilitates the upgrade process of PHP projects. It can help you in 3 major areas:

### 1. Platform Upgrade

Future will test if your project is able to run using the latest platform configuration (e.g. the latest PHP version) and will provide a detailed summary of any encountered issue.

### 2. Dependencies Upgrade

On top of the platform upgrade, it will test if your project is able to run using the latest version of your direct Composer dependencies and will provide a list of blockers that stop your project from upgrading.

### 3. Code Upgrade

Finally, Future will test if your project is compatible with the latest coding style rules and methodologies with the help of Rector.

## Setup

### Prerequisites
* Composer-based Project
* Continuous Integration Pipeline
* Test Suite

### Install
```bash
composer require --dev evozon/future
```

### Configure

TBC

## Running Future

After you have installed and configured Future in your project, run the pipeline and check the output of the `future-proofing` job.

### Output examples
#### Everything is prepared for the upgrade
(output from the pipeline here)
#### This that should be covered before the upgrade
(output from the pipeline here)

## Documentation
TBC

## Contribute
TBC

## Recommendations
Future can be used to test if you can upgrade everything at once: the latest PHP version, the latest versions for your Composer dependencies, and the latest codebase standards. We advise against this since it may lead to massive amounts of changes that are hard to review and test.

We recommend splitting the upgrade process into stages. Start with the PHP version upgrade (one minor version at a time), continue with the dependencies upgrade and finish with the codebase changes. This way you can have smaller PRs that are easier to review and test.

<hr/>

Thanks to [Rector](https://github.com/rectorphp/rector) for providing the tool that does the heavy lifting. Future uses Rector to upgrade the codebase to the latest standards.