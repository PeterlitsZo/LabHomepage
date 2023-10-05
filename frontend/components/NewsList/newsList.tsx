import { useDeleteNews, useListNews } from "../../requests/news";
import List from '../List';

export function NewsList() {
  const listNewsQuery = useListNews();
  const deleteNewsMutation = useDeleteNews();

  return (
    <List
      createItem={{
        url: '/news/create',
        label: 'Create New News',
      }}
      items={{
        whenRender: listNewsQuery.isSuccess,
        refPrefix: '/news/',
        deleteItem(id) {
          deleteNewsMutation.mutate(id);
        },
        getTitle: d => d.title,
        data: listNewsQuery.data?.data.news,
      }}
    />
  )
}
