import axios from 'axios';

export const axiosFetcher = async (url: string) => {
  return axios.get(url).then((res) => res.data);
};

export const axiosPostFetcher = async (url: string, data?: unknown) => {
  return axios.post(url, data).then((res) => res);
};

export const axiosPutFetcher = async (url: string, data?: unknown) => {
  return axios.put(url, data).then((res) => res);
};
