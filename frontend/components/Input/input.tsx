import React from "react";

interface InputProps {
  value?: string,
  onChange?: React.ChangeEventHandler<HTMLInputElement>,
  placeholder?: string,
  type?: string,
  label?: React.ReactNode,
  className?: string,
};

export const Input = (props: InputProps) => {
  const { label, className, ...otherProps } = props;

  return (
    <>
      {label && <label>{label}</label>}
      <input
        className={`w-full p-2 border border-slate-300 rounded focus:outline outline-2 outline-slate-200 ${className ?? ''}`}
        {...otherProps}
      />
    </>
  );
}