import time

import psycopg2
import requests
import json

zabbix_server = 'http://localhost/zabbix'
zabbix_api_user = 'zabbix'
zabbix_api_password = 'Zabbix!@#4'

zabbix_host_name = 'Zabbix server'
zabbix_key_name = 'pg.inserts'

conn = psycopg2.connect(host='localhost', dbname='schooldb',
                        user='mango', password='123123')
cur = conn.cursor()
cur.execute('SELECT COUNT(*) FROM students')
current_count = cur.fetchone()[0]

monitoring_interval = 30
while True:
    cur.execute('SELECT COUNT(*) FROM students')
    new_count = cur.fetchone()[0]
    print(f'{new_count} - количество записей')

    inserts_count = new_count - current_count
    print(f'{inserts_count} - количество записей')
    payload = {
        'request': 'sender data',
        'data': [
            {
                'host': zabbix_host_name,
                'key': zabbix_key_name,
                'value': inserts_count
            }
        ]
    }
    headers = {'Content-type': 'application/json', 'Accept': 'text/plain'}
    response = requests.post(zabbix_server + '/api_jsonrpc.php', data=json.dumps(
        payload), headers=headers, auth=(zabbix_api_user, zabbix_api_password))

    current_count = new_count

    time.sleep(monitoring_interval)
