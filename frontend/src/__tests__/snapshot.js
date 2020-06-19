import React from 'react'
import renderer from 'react-test-renderer'
import HomePage from '../pages/index'
import { shallow, configure } from "enzyme";

import Adapter from 'enzyme-adapter-react-16'

configure({ adapter: new Adapter() })

// const wrapper = shallow(<HomePage />).dive();
it('renders homepage unchanged', () => {
  const tree = renderer.create(<HomePage />).toJSON()
  expect(tree).toMatchSnapshot()
})
