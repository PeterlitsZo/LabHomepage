import { useEffect, useRef } from 'react';

// If click outside of the element controled by ref and the hook is not
// disabled, the call the callback function.
export const useClickOutside = (callback: () => void, disabled: boolean = false) => {
  const ref = useRef<HTMLElement>();

  useEffect(() => {
    const handleClick = (event: Event) => {
      if (
        ref.current
        // We guess the event.target should be a Node.
        && !ref.current.contains(event.target as any)
        && !disabled
      ) {
        callback();
      }
    }

    document.addEventListener('click', handleClick, true);

    return () => {
      document.removeEventListener('click', handleClick, true);
    }
  }, [ref, callback, disabled]);

  return ref;
};