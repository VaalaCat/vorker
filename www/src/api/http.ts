import axios from 'axios'

const instance = axios.create({})

instance.interceptors.request.use((request) => {
  request.headers.Authorization = 'Bearer ' + localStorage.getItem('token')
  return request
})

instance.interceptors.response.use((response) => {
  // console.log(response.headers?.['x-authorization-token'])
  if (response.headers?.['x-authorization-token']) {
    localStorage.setItem('token', response.headers['x-authorization-token'])
  }
  if (!!response.data.code) {
    throw response.data.msg
  }
  return response
})

export default instance
