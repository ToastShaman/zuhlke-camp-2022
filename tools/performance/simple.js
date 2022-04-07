import http from 'k6/http';
import { sleep, check } from 'k6';

export const options = {
    stages: [
        { target: 20, duration: '1m' },
        { target: 15, duration: '1m' },
        { target: 0, duration: '1m' },
    ],
    thresholds: {
        // 90% of requests must finish within 400ms.
        http_req_duration: ['p(90) < 400'],
    },
};

export default function () {
    const host = `${__ENV.ZHOST}`
    const token = `${__ENV.ZTOKEN}`

    const params = {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      };

    const res = http.get(host, params);

    sleep(1);

    check(res, {
        'status is 200': (r) => r.status === 200
    });
}