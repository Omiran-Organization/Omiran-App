import Document, { Head, Main, NextScript } from "next/document";
import React from "react";
import { AppRegistry } from "react-native";
// Force Next-generated DOM elements to fill their parent's height
const normalizeNextElements = `
  body {
    font-family: 'Inter', sans-serif;
  }
`;

export default class MyDocument extends Document {
  static async getInitialProps({ renderPage }) {
    AppRegistry.registerComponent("style-react-native", () => Main);
    const { getStyleElement } = AppRegistry.getApplication(
      "style-react-native"
    );
    const page = renderPage();
    const styles = [
      <style
        type="text/css"
        dangerouslySetInnerHTML={{ __html: normalizeNextElements }}
      />,
      getStyleElement(),
    ];
    return { ...page, styles: React.Children.toArray(styles) };
  }

  render() {
    return (
      <html lang="en">
        <Head>
          <meta name="viewport" content="width=device-width, initial-scale=1" />
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </html>
    );
  }
}
