import { useContext } from 'react';
import ReactMarkdown from 'react-markdown';

import Button from '../Button';

import style from './markdownCard.module.css';
import { useAuthStore } from '../../states/auth';
import useStore from '../../hooks/useStore';

interface MarkdownCardProps {
  title: string;
  content: string;
}

const MarkdownCard = (props: MarkdownCardProps) => {
  const isLoggedIn = useStore(useAuthStore, state => state.isLoggedIn);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 rounded">
        <h1 className="text-4xl font-bold p-8">
          {props.title}
        </h1>
        {isLoggedIn === true && ( // Show tooltip if logged in
          <div className="border-t border-slate-300 px-8 py-4">
            <Button>Update</Button>
          </div>
        )}
        {props.content !== "" && ( // Show content if it is not empty
          <div className="text-lg border-t border-slate-300 p-8">
            <ReactMarkdown className={style.reactMarkdown}>
              {props.content}
            </ReactMarkdown>
          </div>
        )}
      </div>
    </main>
  );
}

export default MarkdownCard;