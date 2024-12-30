package org.util.pack;

import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.*;

public class Pack {

    // PackUint16: 写入一个 16 位无符号整数（Little Endian）
    public static void packUint16(DataOutputStream out, short n) throws IOException {
        out.writeShort(Short.reverseBytes(n)); // Little Endian
    }

    // PackUint32: 写入一个 32 位无符号整数（Little Endian）
    public static void packUint32(DataOutputStream out, int n) throws IOException {
        out.writeInt(Integer.reverseBytes(n)); // Little Endian
    }

    // PackUint64: 写入一个 64 位无符号整数（Little Endian）
    public static void packUint64(DataOutputStream out, long n) throws IOException {
        out.writeLong(Long.reverseBytes(n)); // Little Endian
    }

    // PackString: 写入一个字符串，前面是该字符串的长度（16 位）
    public static void packString(DataOutputStream out, String s) throws IOException {
        packUint16(out, (short) s.length()); // Write length as uint16
        out.write(s.getBytes(StandardCharsets.UTF_8)); // Write the string as bytes
    }

    // PackHexString: 将一个十六进制字符串解码后写入（转成字节后存储）
    public static void packHexString(DataOutputStream out, String hex) throws IOException {
        byte[] bytes = hexStringToByteArray(hex);
        packString(out, new String(bytes, StandardCharsets.UTF_8)); // Write as string
    }

    // hexStringToByteArray: 将十六进制字符串转换为字节数组
    private static byte[] hexStringToByteArray(String hex) {
        int len = hex.length();
        byte[] data = new byte[len / 2];
        for (int i = 0; i < len; i += 2) {
            data[i / 2] = (byte) ((Character.digit(hex.charAt(i), 16) << 4)
                    + Character.digit(hex.charAt(i + 1), 16));
        }
        return data;
    }

    // PackExtra: 写入一个 map，其中 key 为 uint16，value 为字符串
    public static void packExtra(DataOutputStream out, Map<Short, String> extra) throws IOException {
        packUint16(out, (short) extra.size()); // Write the size of the map
        List<Short> keys = new ArrayList<>(extra.keySet());
        Collections.sort(keys); // Sort the keys

        for (Short key : keys) {
            packUint16(out, key); // Write the key (uint16)
            packString(out, extra.get(key)); // Write the value (string)
        }
    }

    // PackMapUint32: 写入一个 map，其中 key 为 uint16，value 为 uint32
    public static void packMapUint32(DataOutputStream out, Map<Short, Integer> extra) throws IOException {
        packUint16(out, (short) extra.size()); // Write the size of the map
        List<Short> keys = new ArrayList<>(extra.keySet());
        Collections.sort(keys); // Sort the keys

        for (Short key : keys) {
            packUint16(out, key); // Write the key (uint16)
            packUint32(out, extra.get(key)); // Write the value (uint32)
        }
    }

//    public static void main(String[] args) {
//        try {
//            // 创建一个 DataOutputStream 用于测试
//            ByteArrayOutputStream byteArrayOutputStream = new ByteArrayOutputStream();
//            DataOutputStream out = new DataOutputStream(byteArrayOutputStream);
//
//            // 示例：使用 packString
//            packString(out, "Hello, World!");
//            // 示例：使用 packUint32
//            packUint32(out, 123456);
//            // 示例：使用 packHexString
//            packHexString(out, "68656c6c6f");
//            // 示例：使用 packExtra
//            Map<Short, String> extra = new HashMap<>();
//            extra.put((short) 1, "value1");
//            extra.put((short) 2, "value2");
//            packExtra(out, extra);
//            // 示例：使用 packMapUint32
//            Map<Short, Integer> mapUint32 = new HashMap<>();
//            mapUint32.put((short) 1, 1000);
//            mapUint32.put((short) 2, 2000);
//            packMapUint32(out, mapUint32);
//
//            // 打印写入的字节
//            byte[] data = byteArrayOutputStream.toByteArray();
//            System.out.println(Arrays.toString(data));
//
//        } catch (IOException e) {
//            e.printStackTrace();
//        }
//    }
}
