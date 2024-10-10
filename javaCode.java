public class javaCode {

    public static void main(String[] args){
        System.out.println("Hello Java");
    }

    // array表示一个长度为n的数组
    // 代码中的array.length就等于n
    private int[] array = new int[100];
    private int count = 0;

    private void insert(int val) {
        if (count == array.length) {
            int sum = 0;
            for (int i = 0; i < array.length; ++i) {
                sum = sum + array[i];
            }
            array[0] = sum;
            count = 1;
        }

        array[count] = val;
        ++count;
    }

    // 全局变量，大小为10的数组array，长度len，下标i。
    int ar[] = new int[10];
    // 全局变量，大小为10的数组array，长度len，下标i。
    int len = 10;
    int i = 0;

    // 往数组中添加一个元素
    void add(int element) {
        if (i >= len) { // 数组空间不够了
            // 重新申请一个2倍大小的数组空间
            int new_array[] = new int[len*2];
            // 把原来array数组中的数据依次copy到new_array
            for (int j = 0; j < len; ++j) {
                new_array[j] = ar[j];
            }
            // new_array复制给array，array现在大小就是2倍len了
            ar = new_array;
            len = 2 * len;
        }
        // 将element放到下标为i的位置，下标i加一
        array[i] = element;
        ++i;
    }
}
