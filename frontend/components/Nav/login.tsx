import { useState, useEffect, useContext } from "react";
import { AuthContext } from "../../contexts/auth";
import { useClickOutside } from "../../hooks/useClickOutside";
import { useLogin } from "../../requests/login";
import Button from "../Button";
import Input from "../Input";

export const LogInButton = () => {
  const [isShowingLogInBox, SetIsShowingLogInBox] = useState(false);

  const auth = useContext(AuthContext);

  const Admin = (
    <div className="rounded p-2 bg-sky-500 leading-none text-white">
      Admin
    </div>
  )

  return (
    <>
      { auth.isLoggedIn
        ? Admin
        : (
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
    </>
  );
}

interface LogInBoxProps {
  isShowingLogInBox: boolean;
  setIsShowingLogInBox: (isShowingLogIn: boolean) => void;
}

const LogInBox = (props: LogInBoxProps) => {
  const [password, setPassword] = useState('');

  const ref = useClickOutside(() => {
    props.setIsShowingLogInBox(false);
  }, !props.isShowingLogInBox);

  const login = useLogin();
  const auth = useContext(AuthContext);
  useEffect(() => {
    if (login.isSuccess) {
      const token = login.data.data.token;
      auth.login(token);
    }
  }, [login, auth]);
  

  return (
    <div
      ref={ref as any}
      className="absolute top-9 right-0 bg-white p-4 w-96 border border-slate-300 rounded"
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
            login.mutate({
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

