import axios from "axios";

export const axiosFetcher = async (url: string) => {
  const accessToken = localStorage.getItem("accessToken");
  return axios
    .get(url, {
      headers: {
        ...(accessToken ? { Authorization: `Bearer ${accessToken}` } : {}),
      },
    })
    .then((res) => res.data);
};

export const axiosPostFetcher = async (url: string, data?: any) => {
  const accessToken = localStorage.getItem("accessToken");
  return axios
    .post(url, data, {
      headers: {
        ...(accessToken ? { Authorization: `Bearer ${accessToken}` } : {}),
      },
    })
    .then((res) => res);
};
