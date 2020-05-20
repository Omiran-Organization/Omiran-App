import React from "react";
import { AppProps } from "next/app";

const Application = ({ Component, pageProps }: AppProps) => (
  <Component {...pageProps} />
);

export default Application;
