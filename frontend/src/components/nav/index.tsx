import * as React from "react";

import Link from "next/link";

const Nav: React.FunctionComponent = () => (
  <div className="fixed flex flex-row items-center h-30 w-screen p-3">
    <Link href="/">
      <a>
        <h1 className="text-blue-500 text-3xl font-bold">Omiran</h1>
      </a>
    </Link>
    <div className="flex-grow" />
    <Link href="/signup">
      <a>
        <button className="btn btn-blue mr-3">Sign Up</button>
      </a>
    </Link>
    <Link href="/login">
      <a>
        <button className="btn btn-blue">Login</button>
      </a>
    </Link>
  </div>
);

export default Nav;
