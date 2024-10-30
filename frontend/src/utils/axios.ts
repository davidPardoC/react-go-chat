import axios from "axios";

export const setAxiosDefaults = () => {
  axios.defaults.baseURL = "http://localhost:5000";
};
