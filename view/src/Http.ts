import axios from "axios";

const HttpClient = axios.create({
    timeout: 30 * 1000,
    headers: {'X-Custom-Header': 'foobar'},
})

export {HttpClient}
