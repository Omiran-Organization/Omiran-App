import * as React from "react";

import PasswordInput from "@/components/input/passwordinput";

const LoginPage: React.FunctionComponent = () => {
  const [emailAddressOrUsername, setEmailAddressOrUsername] = React.useState("");
  const [password, setPassword] = React.useState("");

  return (
    <div className="flex flex-col items-center h-screen w-4/5 md:w-2/5 mx-auto">
      <div className="flex-grow-2" />
      <h1 className="text-blue-500 text-4xl font-bold leading-none">Omiran</h1>
      <h3 className="text-blue-500 text-sm mb-3">
        The Open Source Streaming Platform
      </h3>
      <label
        className="text-blue-500 w-full pl-1"
        htmlFor="email-username-input"
      >
        Email / Username
      </label>
      <input
        className="input w-full mb-3"
        type="text"
        value={emailAddressOrUsername}
        onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          setEmailAddressOrUsername(e.target.value)
        }
        id="email-username-input"
      />
      <label className="text-blue-500 w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <PasswordInput
        containerProps={{ className: "mb-3" }}
        inputProps={{
          value: password,
          onChange: (e: React.ChangeEvent<HTMLInputElement>): void =>
            setPassword(e.target.value),
          id: "password-input",
        }}
      />
      <button className="btn btn-blue self-start">Login</button>
      <div className="flex-grow-3" />
    </div>
  );
};

export default LoginPage;
