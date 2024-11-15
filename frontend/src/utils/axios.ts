import axios from "axios";
import { getCredentials } from "./auth";

export const setAxiosDefaults = () => {
  axios.defaults.baseURL = "http://localhost:5500";
  axios.defaults.headers.common.Authorization = `Bearer ${
    getCredentials().acces_token
  }`;

  axios.interceptors.response.use(
    (response) => response,
    (error) => {
      if (error.response?.status === 401) {
        window.location.href = "/login";
      }
      return Promise.reject(error);
    }
  );
};
