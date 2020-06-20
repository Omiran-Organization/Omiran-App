import React from 'react'
import renderer from 'react-test-renderer'
import Link from '../components/Link.react'; 
import HomePage from '../pages/index'
import { shallow, configure } from "enzyme";

import Adapter from 'enzyme-adapter-react-16'

configure({ adapter: new Adapter() })

// const wrapper = shallow(<HomePage />).dive();
// it('renders homepage unchanged', () => {
//   const tree = renderer.create(<HomePage />).toJSON()
//   expect(tree).toMatchSnapshot()
// })
it('renders correctly', () => {
  const tree = renderer
    .create(<Link page="http://www.facebook.com">Facebook</Link>)
    .toJSON();
  expect(tree).toMatchSnapshot();
});