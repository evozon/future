# Configuring Future for GitLab

This page provides a set of instructions that will help you configure Future for GitLab.

1. Open the `.gitlab-ci.yml` file and create the `future` stage
```
stages:
    - ...
    - future
```

2. Copy the job defined in the `future-proofing.yaml` file ([link](https://github.com/evozon/future/blob/master/future-proofing.yaml)) into `.gitlab-ci.yml`

```
future-proofing:
    ...
```

3. Configure the platform you target for the upgrade by defining the base image for the `future-proofing` job. If you project has a Dockerfile, use that instead and update its `FROM` instruction to match the targeted platform.

```
future-proofing:
    ...
    image: php:fpm-alpine
    ...
```

4. Define the `services` required for running your test suite. This section will be similar to the job that executes your project's test suite.

```
future-proofing:
    ...
    services:
        - mysql:latest
    ...
```

5. Define the environment `variables` required for running your test suite. This section will be similar to the job that executes your project's test suite.

```
future-proofing:
    ...
    variables:
        MYSQL_DATABASE: $TEST_MYSQL_DATABASE
        MYSQL_ROOT_PASSWORD: $TEST_MYSQL_ROOT_PASSWORD
        ...
    ...
```

6. Run any scripts and installation steps required for executing your project's test suite in the `before_script` section of the `future-proofing` job (install Composer dependencies, install PHP extensions, install Xdebug, run migrations, load fixtures, etc.)

```
future-proofing:
    ...
    before_script:
        ...
        - docker-php-ext-install ldap pdo_mysql
        - curl --silent --show-error "https://getcomposer.org/installer" | php -- --install-dir=/usr/local/bin --filename=composer
        - composer install --no-interaction --no-scripts --prefer-dist
        - bin/console --env=test doctrine:schema:create
        - bin/console --env=test doctrine:fixtures:load --no-interaction
        ...
    ...
```

7. Configure the upgrade process in the `script` section. By default, Future will bump PHP to the latest version, will bump your direct Composer dependencies to their latest version, and will apply Rector to your code with a handpicked set of rectors. Everything is fully configurable, and you can keep whatever you need for your upgrade

```
future-proofing:
    ...
    script:
        - vendor/bin/future bump-php latest
        - vendor/bin/future bump-deps
        - composer update -W --no-interaction --no-scripts
        ...
    ...
```

8. Configure your test suite to run at the end of the `script` section

```
future-proofing:
    ...
    script:
        ...
        - vendor/bin/phpunit --no-interaction --colors=never
    ...
```