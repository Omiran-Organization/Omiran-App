import React from "react";

import { AppProps } from "next/app";

import "../public/css/style.css";

const Application = ({
  Component,
  pageProps,
}: AppProps): React.ReactElement => <Component {...pageProps} />;

export default Application;