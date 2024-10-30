import { Credentials } from "@/entities/credentials";
import { User } from "@/entities/user";
import axios from "axios";
import { LoginDto, SignupDto } from "../dtos/signup-dto";

export const signupUser = async (body: SignupDto): Promise<User> => {
  const { data } = await axios.post<User>("/v1/auth/signup", body);
  return data;
};

export const loginUser = async (body: LoginDto): Promise<Credentials> => {
  const { data } = await axios.post<Credentials>("/v1/auth/login", body);
  return data;
};
