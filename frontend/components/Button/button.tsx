import React, { FC, ReactNode, useState } from "react";
import { useClickOutside } from "../../hooks/useClickOutside";

interface ButtonProps {
  children?: ReactNode;
  className?: string;
  disabled?: boolean;
  onClick?: () => void;
  onClickOutside?: () => void;
}

export const Button: FC<ButtonProps> = (props) => {
  return (
    <button
      className={
        `inline-flex h-8 px-4 cursor-pointer rounded bg-sky-600 hover:bg-sky-700 text-white items-center justify-center content-center ${
          props.className ?? ''
        } ${
          props.disabled && 'bg-slate-500 hover:bg-slate-500' 
        }`
      }
      disabled={props.disabled}
      onClick={() => {
        props.onClick && props.onClick();
      }}
    >
      {props.children}
    </button>
  );
};