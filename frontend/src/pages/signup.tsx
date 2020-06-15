import * as React from "react";

import Head from "next/head";

import PasswordInput from "@/components/input/passwordinput";

const emailAddressRegex = /^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$/;

const SignupPage: React.FunctionComponent = () => {
  const [emailAddress, setEmailAddress] = React.useState("");
  const [username, setUsername, setDescription, description] = React.useState("");

  const [password, setPassword] = React.useState("");
  const [passwordConfirmation, setPasswordConfirmation] = React.useState("");

  // Form validation
  const isEmailAddressValid = emailAddressRegex.test(emailAddress);
  const isUsernameValid = username.length > 0;

  const isPasswordValid = password.length > 0;
  const arePasswordsMatching = password == passwordConfirmation;

  const isFormComplete =
    isEmailAddressValid &&
    isUsernameValid &&
    isPasswordValid &&
    arePasswordsMatching;

  return (
    <div className="main flex flex-col items-center w-4/5 md:w-2/5 mx-auto">
      <Head>
        <title>Signup - Omiran</title>
      </Head>
      <div className="flex-grow-2" />
      <h1 className="text-4xl font-bold leading-none">Omiran</h1>
      <h3 className="text-sm mb-2">The Open Source Streaming Platform</h3>
      <label className="w-full pl-1" htmlFor="email-input" name="email">
        Email
      </label>
      <input
        className={`input ${
          !isEmailAddressValid && emailAddress.length > 0 ? "input-error" : ""
        } w-full mb-2`}
        type="email"
        value={emailAddress}
        onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          setEmailAddress(e.target.value)
        }
        id="email-input"
      />
      <label className="w-full pl-1" htmlFor="username-input">
        Username
      </label>
      <input
        className="input w-full mb-2"
        type="text"
        value={username}
        onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          setUsername(e.target.value)
        }
        id="username-input"
        name="username"
      />
      <label className="w-full pl-1" htmlFor="description-input">
        Description
      </label>
      <input
        className="input w-full mb-2"
        type="text"
        value={description}
        onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          setDescription(e.target.value)
        }
        id="description-input"
        name="description"     
      />
      <label className="w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <PasswordInput
        containerProps={{ className: "mb-2" }}
        inputProps={{
          value: password,
          onChange: (e: React.ChangeEvent<HTMLInputElement>): void =>
            setPassword(e.target.value),
          id: "password-input",
        }}
        name="password"
      />
      <label className="w-full pl-1" htmlFor="password-confirmation-input">
        Confirm Password
      </label>
      <PasswordInput
        containerProps={{ className: "mb-3" }}
        inputProps={{
          value: passwordConfirmation,
          onChange: (e: React.ChangeEvent<HTMLInputElement>): void =>
            setPasswordConfirmation(e.target.value),
          id: "password-confirmation-input",
        }}
      />
      <div className="flex flex-row items-center w-full">
        <button
          className={`btn ${isFormComplete ? "btn-orange" : "btn-disabled"} mr-4`}
        >
          Sign Up
        </button>
        <span className="text-red-500">
          {!arePasswordsMatching ? "Passwords do not match." : ""}
        </span>
      </div>
      <div className="flex-grow-3" />
    </div>
  );
};

export default SignupPage;
