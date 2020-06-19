import React from "react";

import Document, { Head, Main, NextScript } from "next/document";

import Nav from "../components/Nav";

export default class MyDocument extends Document {
  render(): React.ReactElement {
    return (
      <html lang="en">
        <Head>
          <meta name="viewport" content="width=device-width, initial-scale=1" />
        </Head>
        <body>
          <Nav />
          <Main />
          <NextScript />
        </body>
      </html>
    );
  }
}
