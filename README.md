#Moiseev test work

For run webserver and tests need install docker and make utility.

1. Create directory for clone repository:
```
mkdir /tmp/moisieiev_test
```
2. Clone repository:
```
cd /tmp/moisieiev_test
git clone git@github.com:Ifaliuk/yalantis_test.git .
```
3. For run unit tests:
```
make run_tests
```
4. For run webserver:
```
make dc_up
```
and open in browser: http://localhost:8110 or change port in docker-compose.yaml file.
Application logs directory location: logs

5. For stop webserver:
```
make dc_down
```