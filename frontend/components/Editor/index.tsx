import { useState, ReactNode, FC } from "react";
import Button from '../Button'
import { Edit, View } from 'lucide-react';

interface EditorProps {
  submit: (titleContent: { title: string, content: string }) => void;
  titlePlaceholder: string;
  contentPlaceholder: string;
}

interface TabItemProps {
  children: ReactNode;
  focused?: boolean;
}

const TabItem: FC<TabItemProps> = (props) => {
  return (
    <button
      className={`inline-flex items-center p-2 rounded ${
        props.focused ? "bg-sky-600 text-white" : "bg-white text-black"
      }`}
    >
      {props.children}
    </button>
  );
}

const Editor: FC<EditorProps> = (props) => {
  const [title, setTitle] = useState('');
  const [content, setContent] = useState('');

  return (
    <main className="max-w-7xl m-auto my-24 border border-slate-300 rounded">
      <input
        placeholder={props.titlePlaceholder}
        className="w-full p-8 text-4xl outline-none"
        onChange={(e) => setTitle(e.target.value)}
      />
      <div className="flex items-center bg-slate-100 px-8 py-4 h-20">
        <div className="bg-white rounded p-2">
          <TabItem focused>
            <Edit style={{ width: '1rem', height: '1rem' }} />&nbsp;Edit
          </TabItem>
          <TabItem>
            <View style={{ width: '1rem', height: '1rem' }} />&nbsp;Preview
          </TabItem>
        </div>
        <span className="flex-1" />
        <Button className="h-12 w-24 shadow" onClick={() => {
          props.submit({ title, content });
        }}>
          Publish
        </Button>
      </div>
      <textarea
        rows={20}
        placeholder={props.contentPlaceholder}
        className="w-full p-8 text-xl outline-none"
        onChange={(e) => setContent(e.target.value)}
      />
    </main>
  )
};

export default Editor;