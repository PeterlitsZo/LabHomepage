import { useContext } from 'react';
import ReactMarkdown from 'react-markdown';

import { AuthContext } from '../../contexts/auth';
import Button from '../Button';

import style from './markdownCard.module.css';

interface MarkdownCardProps {
  title: string;
  content: string;
}

const MarkdownCard = (props: MarkdownCardProps) => {
  const auth = useContext(AuthContext);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="border border-slate-300 rounded">
        <h1 className="text-4xl font-bold p-8">
          {props.title}
        </h1>
        {auth.isLoggedIn === true && ( // Show tooltip if logged in
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