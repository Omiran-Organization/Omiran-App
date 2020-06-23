import React from 'react'

import Link from 'next/link'

export default function Nav() {
  return (
    <header>
      <div className="relative flex bg-orange-500 flex-row items-center top-0 h-20 w-screen p-3">
        <Link href="/">
          <a>
            <h1 className="text-black text-3xl font-bold">Omiran</h1>
          </a>
        </Link>
        <div className="flex-grow" />
        <Link href="/signup">
          <a>
            <button className="btn btn-orange btn-outlined mx-1">Sign Up</button>
          </a>
        </Link>
        <Link href="/login">
          <a>
            <button className="btn btn-orange btn-outlined mx-1">Login</button>
          </a>
        </Link>
      </div>
    </header>
  )
}
