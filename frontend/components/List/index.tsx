import Link from "next/link";
import { useAuthStore } from "../../states/auth";
import { PlusCircle, Trash2 } from "lucide-react";
import classNames from "classnames";
import useStore from "../../hooks/useStore";

interface CreateItem {
  url: string;
  label: string;
}

interface ItemData {
  id: string;
}

interface Items<T extends ItemData> {
  whenRender: boolean;
  refPrefix: string;
  deleteItem(id: string): void;
  getTitle(d: T): string;
  data?: T[];
};

interface ListProps<T extends ItemData> {
  createItem: CreateItem;
  items: Items<T>;
}

export function List<T extends ItemData>(props: ListProps<T>) {
  const isLoggedIn = useStore(useAuthStore, state => state.isLoggedIn);

  return (
    <main className="max-w-3xl m-auto py-24">
      <div className="text-xl p-4 rounded border border-slate-300">
        {isLoggedIn && (
          <Link
            href={props.createItem.url}
            className={classNames(
              "flex content-center items-center gap-2 p-4 h-14 rounded",
              "text-slate-500 hover:text-sky-900 cursor-pointer hover:bg-slate-100",
            )}
          >
            <PlusCircle />
            Create New News
          </Link>
        )}
        {props.items.whenRender && (() => {
          let data = props.items.data!;
          return [...data].reverse().map(d => (
            <Link
              id={d.id}
              href={props.items.refPrefix + d.id}
              className="flex content-center items-center p-4 h-14 rounded hover:bg-slate-100 cursor-pointer"
            >
              <span>
                {props.items.getTitle(d)}
              </span>
              <span className="flex-1" />
              {isLoggedIn && (
                <Trash2
                  onClick={(e: MouseEvent) => {
                    e.stopPropagation();
                    props.items.deleteItem(d.id);
                  }}
                  className="hover:bg-slate-200 p-2 h-10 w-10 rounded-full text-slate-500"
                />
              )}
            </Link>
          ))
        })()}
      </div>
    </main>
  )
};

export default List;