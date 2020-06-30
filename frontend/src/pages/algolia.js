import algoliasearch from 'algoliasearch/lite';
import React, { Component } from 'react';
import {
  InstantSearch,
  Hits,
  SearchBox,
  Pagination,
  Highlight,
  ClearRefinements,
  RefinementList,
  Configure,
} from 'react-instantsearch-dom';
import PropTypes from 'prop-types';

import Head from "next/head";




const searchClient = algoliasearch(
    process.env.NEXT_PUBLIC_ALGOLIA_APP_ID,
  process.env.NEXT_PUBLIC_ALGOLIA_SEARCH_KEY
);


class App extends Component {
  render() {
    return (
        
       
      <div className="main flex flex-col">
      <Head>
        <title> Omiran</title>
      </Head>
         <h1>React InstantSearch e-commerce demo</h1>
         <InstantSearch indexName="demo_ecommerce" searchClient={searchClient} >
           <div className="left-panel">
             <ClearRefinements />
             <h2>Brands</h2>
             {/* <RefinementList attribute="brand" /> */}
             {/* <Configure hitsPerPage={8} /> */}
           </div>

           <div className="right-panel py-10">
             <SearchBox className="input w-64 mb-3"/>
             <Hits hitComponent={Hit} />
             <Pagination />
           </div>
         </InstantSearch>


        <div className="flex-grow" />
  </div>

  )}
}





function Hit(props) {
  return (
    <div>
      <img src={props.hit.image} align="left" alt={props.hit.name} />
      <div className="hit-name">
        <Highlight attribute="name" hit={props.hit} />
      </div>
      <div className="hit-description">
        <Highlight attribute="description" hit={props.hit} />
      </div>
      <div className="hit-price">${props.hit.price}</div>
    </div>
  );
}

Hit.propTypes = {
  hit: PropTypes.object.isRequired,
};

export default App;
