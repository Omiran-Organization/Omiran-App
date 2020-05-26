import * as React from "react";

const LoginPage: React.FunctionComponent = () => {
  const [emailOrUsername, setEmailOrUsername] = React.useState("");

  const [password, setPassword] = React.useState("");
  const [passwordShown, setPasswordShown] = React.useState(false);

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
        value={emailOrUsername}
        onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
          setEmailOrUsername(e.target.value)
        }
        id="email-username-input"
      />
      <label className="text-blue-500 w-full pl-1" htmlFor="password-input">
        Password
      </label>
      <div className="input flex flex-row w-full mb-3">
        <input
          className="w-full"
          type={passwordShown ? "text" : "password"}
          value={password}
          onChange={(e: React.ChangeEvent<HTMLInputElement>): void =>
            setPassword(e.target.value)
          }
          id="password-input"
        />
        <div
          className="text-gray-700 hover:text-gray-800"
          onClick={(): void => setPasswordShown(!passwordShown)}
        >
          {passwordShown ? (
            // prettier-ignore
            <svg className="h-6 s-6 my-auto" fill="currentColor" viewBox="0 0 20 20"><path d="M10 12a2 2 0 100-4 2 2 0 000 4z"></path><path d="M.458 10C1.732 5.943 5.522 3 10 3s8.268 2.943 9.542 7c-1.274 4.057-5.064 7-9.542 7S1.732 14.057.458 10zM14 10a4 4 0 11-8 0 4 4 0 018 0z" clipRule="evenodd" fillRule="evenodd"></path></svg>
          ) : (
            // prettier-ignore
            <svg className="h-6 s-6 my-auto" fill="currentColor" viewBox="0 0 20 20"><path d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z" clipRule="evenodd" fillRule="evenodd"></path><path d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z"></path></svg>
          )}
        </div>
      </div>
      <button className="btn btn-blue self-start">Login</button>
      <div className="flex-grow-3" />
    </div>
  );
};

export default LoginPage;
