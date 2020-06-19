import Link from 'next/link'
import { useRouter } from 'next/router'

export default function Nav() {
  // const { pathname } = useRouter()

  return (
    <header>
      <div className="fixed flex bg-orange-500 flex-row items-center top-0 h-20 w-screen p-3">
    <Link href="/">
      <a>
        <h1 className="text-black text-3xl font-bold">Omiran</h1>
      </a>
    </Link>
    <div className="flex-grow" />
    <Link href="/signup">
      <a>
        <button className="btn btn-orange btn-outlined mr-3">Sign Up</button>
      </a>
    </Link>
    <Link href="/login">
      <a>
        <button className="btn btn-orange btn-outlined">Login</button>
      </a>
    </Link>
    <Link href="/profile">
      <a>
        <button className="btn btn-orange btn-outlined">Profile</button>
      </a>
    </Link>
    {/* <Link href="/apollo">
        <a className={pathname === '/apollo' ? 'is-active' : ''}>Apollo</a>
      </Link>
      <Link href="/redux">
        <a className={pathname === '/redux' ? 'is-active' : ''}>Redux</a>
      </Link> */}
      
    </div>
  </header>
    
  )
}
