import Document, { Head, Main, NextScript } from "next/document";
import React from "react";

import Nav from "@/components/nav";

export default class MyDocument extends Document {
  render() {
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
