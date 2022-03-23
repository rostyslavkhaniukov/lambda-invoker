FROM bref/php-80-console

COPY --from=bref/extra-yaml-php-80:0.11.19 /opt/bref-extra /opt/bref-extra
COPY --from=bref/extra-yaml-php-80:0.11.19 /opt/bref/ /opt/bref

COPY . /var/task

CMD ["artisan"]
