import Link from 'next/link';
import { useRouter } from 'next/router';
import React, { ReactNode } from 'react';
import classNames from 'classnames';

import { AuthButton } from './AuthButton';

export const Nav = () => {
  return (
    <nav className="relative flex border-b border-slate-300 justify-center bg-white px-4 lg:px-8 py-4">
      <div
        className={classNames(
          "flex-auto shrink-0",
          "max-w-7xl h-7",
          "flex items-center content-center",
        )}
      >
        <Link href="/" className="font-bold text-xl tracking-wide leading-none py-2 hover:text-slate-500">
          Lab Homepage
        </Link>
        <NavLink href="/news">News</NavLink>
        <NavLink href="/papers">Papers</NavLink>
        <NavLink href="/people">People</NavLink>
        <NavLink href="/resources">Resources</NavLink>
        <span className="flex-1" />
        <AuthButton />
      </div>
    </nav>
  )
};

interface NavLinkProps {
  children: ReactNode;
  href: string;
}

const NavLink = (props: NavLinkProps) => {
  const router = useRouter();
  const isCurrent = router.pathname.startsWith(props.href);

  return (
    <Link
      href={props.href}
      className={`text-lg ml-4 p-2 ${isCurrent ? 'font-bold text-sky-900' : 'text-slate-400 hover:text-black'} cursor-pointer rounded hover:bg-slate-100 leading-none`}
    >
      {props.children}
    </Link>
  )
}

