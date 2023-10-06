import React, { useEffect } from 'react';
import * as Avatar from '@radix-ui/react-avatar';
import * as DropdownMenu from '@radix-ui/react-dropdown-menu';

import { useAuthStore } from '../../../states/auth';
import useStore from '../../../hooks/useStore';

export function AdminAvatar() {
  const logout = useAuthStore(state => state.logout);

  return (
    <DropdownMenu.Root>
      <DropdownMenu.Trigger asChild>
        <Avatar.Root>
          <Avatar.Fallback className="rounded p-2 bg-sky-500 leading-none text-white">
            ADMIN
          </Avatar.Fallback>
        </Avatar.Root>
      </DropdownMenu.Trigger>
      <DropdownMenu.Portal>
        <DropdownMenu.Content
          className="rounded py-2 bg-white border border-slate-300 w-40"
          sideOffset={16}
          align="end"
        >
          <DropdownMenu.Item asChild onClick={logout}>
            <div className="px-4 py-2 hover:bg-slate-100 cursor-pointer">
              Log out
            </div>
          </DropdownMenu.Item>
        </DropdownMenu.Content>
      </DropdownMenu.Portal>
    </DropdownMenu.Root>
  )
}