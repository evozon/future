# Configuring Future for GitHub

This page provides a set of instructions that will help you configure Future for GitHub.

1. Create the `future-proofing` workflow from the [template](https://github.com/evozon/future/blob/master/ci/github/future-proofing.yaml) provided in this library. We assume the presence of the `.github/workflows/` directory since having a CI pipeline is one of the prerequisites of using Future

```
cp vendor/evozon/future/ci/github/future-proofing.yaml .github/workflows/
```

2. Configure the platform you target for the upgrade by setting a valid value for the `runs-on` key

```
...
jobs:
    future-proofing:
        name: Future Proofing
        runs-on: ubuntu-latest
        ...
```

3. Define the `services` required for running your test suite. If no service is required, you can omit this step. This section will be similar to the job that executes your project's test suite

```
...
jobs:
    future-proofing:
        ...
        services:
            mysqldb:
                image: mysql:latest
                ...
```

4. Run any custom scripts and installation steps required for executing your project's test suite (install Composer dependencies, install PHP extensions, install Xdebug, run migrations, load fixtures, etc.). If no custom logic is required, you can omit this step

```
...
jobs:
    future-proofing:
        ...       
        - name: Custom installation script
          run: ...
```

5. Configure the upgrade process. By default, Future will bump PHP to the platform's version, will bump your direct Composer dependencies to their latest version, and will apply Rector to your code with a handpicked set of rectors. Everything is fully configurable, and you can keep whatever you need for your upgrade

```
...
jobs:
    future-proofing:
        ...
        - name: Bump PHP
          run: vendor/bin/future bump-php

        - name: Bump dependencies
          run: vendor/bin/future bump-deps

        - name: Update dependencies
          run: composer update -W --no-interaction --no-scripts

        - name: Update codebase
          run: ...
```

6. Configure your test suite to run at the end of the `future-proofing` job

```
...
jobs:
    future-proofing:
        ...
        - name: Execute PHPUnit tests
          run: vendor/bin/phpunit --no-interaction --colors=never
```

_Note: Please make sure you provide all the necessary environment variables under the `env` key of each step._
