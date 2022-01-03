#include "../lib/libdga.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

int main(int argc, char **argv)
{
    char* key = GenerateRandomKey(10);
    char* format = "yyyyMMddHHmm";
    time_t now = time(NULL);
    struct tm *unixtime = gmtime(&now);
    int* count = 10;
    char* suffix = ".com";

    if (argc == 1)
    {
        printf("Usage: dga [options]\n");
        printf("Options:\n");
        printf("  -k\tKey, Base32 String\n");
        printf("  -t\tUnix-Time Seconds, Default: Now\n");
        printf("  -c\tGenerator Count, Default: 10\n");
        printf("  -f\tDate Format Pattern, Default: yyyyMMddHHmm\n");
        printf("  -s\tDomain Suffix, Default: .com\n");
        printf("Example:\n");
        printf("  dga -k %s -t %d\n", key, mktime(unixtime));
        return 0;
    }

    for (int i = 1; i < argc; i += 2)
    {
        if (strcmp(argv[i], "-k") == 0)
            key = argv[i + 1];
        else if (strcmp(argv[i], "-t") == 0) {
            now = (time_t)atoi(argv[i + 1]);
            unixtime = gmtime(&now);
        }
        else if (strcmp(argv[i], "-c") == 0)
            count = atoi(argv[i + 1]);
        else if (strcmp(argv[i], "-f") == 0)
            format = argv[i + 1];
        else if (strcmp(argv[i], "-s") == 0)
            suffix = argv[i + 1];
    }

    if (count < 1)
        count = 1;

    char **domains = GenerateDomain(key, now, format, count);

    printf("Key: %s\n", key);
    printf("Unix-Time Seconds: %d\n", now);
    printf("DateTime: %s", asctime(unixtime));
    printf("Format: %s\n", format);
    printf("Count: %d\n", count);
    printf("Suffix: %s\n", suffix);
    printf("Domains:");

    for (int i = 0; i < count; i++)
    {
        printf(" %s%s", domains[i], suffix);
    }
    printf("\n");

    return 0;
}