import * as React from "react";

import Head from "next/head";
import { useRouter } from "next/router";

import PasswordInput from "@/components/input/passwordinput";

const LoginPage: React.FunctionComponent = () => {
  const router = useRouter();

  const [emailAddressOrUsername, setEmailAddressOrUsername] = React.useState(
    ""
  );
  const [password, setPassword] = React.useState("");

  return (
    <div className="main flex flex-col items-center w-4/5 md:w-2/5 mx-auto">
      <Head>
        <title>Login - Omiran</title>
      </Head>
      <div className="flex-grow-2" />
      <h1 className="text-4xl font-bold leading-none">Omiran</h1>
      <h3 className="text-sm mb-3">The Open Source Streaming Platform</h3>
      <label className="w-full pl-1" htmlFor="email-username-input">
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
      <label className="w-full pl-1" htmlFor="password-input">
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
      <button
        className="btn btn-orange self-start"
        onClick={(): Promise<boolean> => router.push("/profile")}
      >
        Login
      </button>
      <div className="flex-grow-3" />
    </div>
  );
};

export default LoginPage;
