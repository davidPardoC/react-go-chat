import axios from "axios";
import { User } from "../entities/user.entity";

export const getAllUsers = async ():Promise<User[]> => {
  const { data } = await axios.get("/v1/users/");
  return data;
};
