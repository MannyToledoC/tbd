import styles from "./page.module.css";
import Test from "@/Components/Test/Test";

export default function Home() {
  return (
    <main className={styles.main}>
      <div className={styles.description}>
        <Test />
      </div>
    </main>
  );
}
