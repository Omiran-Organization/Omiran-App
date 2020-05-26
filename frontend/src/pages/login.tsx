import * as React from "react";

const LoginPage: React.FunctionComponent = () => (
  <div className="flex flex-col justify-center items-center h-screen w-4/5 md:w-2/5 mx-auto">
    <h1 className="text-blue-500 text-4xl font-bold text-center leading-none">
      Omiran
    </h1>
    <h3 className="text-blue-500 text-sm mb-3">
      The Open Source Streaming Platform
    </h3>
    <label className="text-blue-500 w-full pl-1" htmlFor="email-username-input">
      Email / Username
    </label>
    <input
      className="input w-full mb-3"
      type="text"
      id="email-username-input"
    />
    <label className="text-blue-500 w-full pl-1" htmlFor="password-input">
      Password
    </label>
    <input className="input w-full mb-3" type="password" id="password-input" />
    <button className="btn btn-blue self-start">Login</button>
  </div>
);

export default LoginPage;
