import { useState, useEffect, useContext } from "react";

import { useClickOutside } from "../../../hooks/useClickOutside";
import { useLogin } from "../../../requests/login";
import Button from "../../Button";
import Input from "../../Input";
import { useAuthStore } from "../../../states/auth";
import useStore from "../../../hooks/useStore";
import { AdminAvatar } from "./AdminAvatar";

export function AuthButton() {
  const [isShowingLogInBox, SetIsShowingLogInBox] = useState(false);

  const isLoggedIn = useStore(useAuthStore, state => state.isLoggedIn);

  if (isLoggedIn) {
    return (
      <AdminAvatar />
    )
  } else {
    return (
      <div className="relative">
        <Button
          onClick={() => SetIsShowingLogInBox(true)}
          disabled={isShowingLogInBox}
          className="-my-1"
        >
          Log in
        </Button>
        {isShowingLogInBox && (
          <LogInBox
            isShowingLogInBox={isShowingLogInBox}
            setIsShowingLogInBox={SetIsShowingLogInBox}
          />
        )}
      </div>
    )
  }
}

interface LogInBoxProps {
  isShowingLogInBox: boolean;
  setIsShowingLogInBox: (isShowingLogIn: boolean) => void;
}

function LogInBox(props: LogInBoxProps) {
  const [password, setPassword] = useState('');

  const ref = useClickOutside(() => {
    props.setIsShowingLogInBox(false);
  }, !props.isShowingLogInBox);

  const { isSuccess: loginSuccess, data: loginData, mutate: loginMutate } = useLogin();
  const login = useAuthStore(state => state.login);

  useEffect(() => {
    console.log('login: ', login);
  }, [login]);

  useEffect(() => {
    if (loginSuccess && login) {
      const token = loginData.data.token;
      console.log('call login: ', login);
      login(token);
    }
  }, [loginSuccess, login]);
  

  return (
    <div
      ref={ref as any}
      className="absolute top-9 right-0 z-10 bg-white p-4 w-96 border border-slate-300 rounded"
    >
      <div className="font-bold text-xl mb-4">Welcome back</div>
      <Input
        placeholder="Please enter the admin's password."
        type="password"
        onChange={(e) => setPassword(e.target.value)}
      />
      <div className="flex flex-row-reverse mt-4">
        <Button
          onClick={() => {
            loginMutate({
              username: 'admin',
              password: password,
            });
          }}
        >
          Log in
        </Button>
      </div>
    </div>
  )
}

