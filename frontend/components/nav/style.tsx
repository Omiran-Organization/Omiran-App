import { StyleSheet } from "react-native-web";

export default StyleSheet.create({
  header: {
    width: "100%",
    paddingHorizontal: 20,
    paddingVertical: 10,
    display: "flex",
    flexWrap: "wrap",
    position: "absolute",
    flexDirection: "row",
    boxShadow: "0 3px 6px rgba(0, 0, 0, 0.2)",
    alignItems: "center",
  },
  col: {
    flex: 1,
  },
  colLarge: {
    flexBasis: "60%",
  },
  logo: {
    maxWidth: 250,
    height: 40,
    flexGrow: 0,
  },
  search: {
    width: "100%",
    maxWidth: 400,
    height: 40,
  },
  searchInput: {
    backgroundColor: "#e0dfdc",
  },
  menu: {
    display: "flex",
    flexDirection: "row",
    justifyContent: "flex-end",
  },
  menuItem: {
    color: "#000",
    fontSize: 20,
  },
});
