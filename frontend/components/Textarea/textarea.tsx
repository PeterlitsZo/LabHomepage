import React from "react";

interface InputProps {
  value?: string,
  onChange?: React.ChangeEventHandler<HTMLTextAreaElement>,
  placeholder?: string,
  type?: string,
  label?: React.ReactNode,
  rows?: number,
};

export const Textarea = (props: InputProps) => {
  const { label, ...otherProps } = props;

  return (
    <>
      {label && <label>{label}</label>}
      <textarea
        className="w-full p-2 border border-slate-300 rounded focus:outline outline-2 outline-slate-200"
        {...otherProps}
      />
    </>
  );
}