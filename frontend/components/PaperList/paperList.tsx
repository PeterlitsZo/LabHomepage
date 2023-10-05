import { useDeletePaper, useListPaper } from "../../requests/papers"
import List from "../List";

export const PaperList = () => {
  const listPaperQuery = useListPaper();
  const deletePaperMutation = useDeletePaper();

  return (
    <List
      createItem={{
        url: '/papers/create',
        label: 'Create New Paper',
      }}
      items={{
        whenRender: listPaperQuery.isSuccess,
        refPrefix: '/pagers/',
        deleteItem(id) {
          deletePaperMutation.mutate(id);
        },
        getTitle: d => d.title,
        data: listPaperQuery.data?.data.papers,
      }}
    />
  );
}