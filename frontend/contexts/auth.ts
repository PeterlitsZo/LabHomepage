import { createContext } from "react";

interface AuthContextValue {
  token: string;
  isLoggedIn: boolean;
  login: (token: string) => void;
}

export const AuthContext = createContext({} as AuthContextValue);