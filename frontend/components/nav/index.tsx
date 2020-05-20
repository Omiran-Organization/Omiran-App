import React, { useState } from "react";
import Link from "next/link";
import { View, Text, Image, TextInput } from "react-native-web";

import styles from "./style";

const Nav: React.FunctionComponent = () => {
  const [searchValue, setSearchValue] = useState("");

  return (
    <View style={styles.header}>
      <View style={styles.logo}>
        <Image
          source={{
            uri: "https://via.placeholder.com/150/40",
            width: 200,
            height: 40,
          }}
        />
      </View>
      <View style={styles.search}>
        <TextInput
          style={styles.searchInput}
          onChangeText={(text) => setSearchValue(text)}
          value={searchValue}
        />
      </View>
      <View style={[styles.colLarge, styles.menu]}>
        <View style={styles.menuItem}>
          <Link href="/browse">
            <Text>Browse</Text>
          </Link>
        </View>
        <View style={styles.menuItem}>
          <Link href="/browse">
            <Text>Browse</Text>
          </Link>
        </View>
        <View style={styles.menuItem}>
          <Link href="/browse">
            <Text>Browse</Text>
          </Link>
        </View>
        <View style={styles.menuItem}>
          <Link href="/browse">
            <Text>Browse</Text>
          </Link>
        </View>
      </View>
    </View>
  );
};

export default Nav;
