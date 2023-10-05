import { create } from 'zustand';
import { persist, createJSONStorage } from 'zustand/middleware'

interface AuthStore {
  token: string;
  isLoggedIn: boolean;
  login(token: string): void;
  logout(): void;
}

export const useAuthStore = create<AuthStore>()(
  persist(
    set => ({
      token: '',
      isLoggedIn: false,
      login(token) {
        set(() => ({
          token,
          isLoggedIn: true,
        }))
      },
      logout() {
        set(() => ({
          token: '',
          isLoggedIn: false,
        }))
      }
    }),
    {
      name: 'auth',
      storage: createJSONStorage(() => sessionStorage),
    }
  )
);