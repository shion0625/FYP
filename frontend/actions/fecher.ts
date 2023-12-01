import axios from "axios";

export const axiosFetcher = (url: string) => axios.get(url).then((res) => res.data);

export const axiosPostFetcher = (url: string, data: any) =>
  axios.post(url, data).then((res) => res.data);
