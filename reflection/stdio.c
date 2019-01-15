  ...
  for (cnt = 0; cnt < nargs; ++cnt)
    switch (args_type[cnt])
      {
#define T(tag, mem, type)                                \
        case tag:                                        \
          args_value[cnt].mem = va_arg (*ap_savep, type); \
          break
        T (PA_WCHAR, pa_wchar, wint_t);
      case PA_CHAR:                                /* Promoted.  */
      case PA_INT|PA_FLAG_SHORT:                /* Promoted.  */
#if LONG_MAX == INT_MAX
      case PA_INT|PA_FLAG_LONG:
#endif
        T (PA_INT, pa_int, int);
      ...
      case PA_FLOAT:                                /* Promoted.  */
        T (PA_DOUBLE, pa_double, double);
      case PA_STRING:                                /* All pointers are the same */
      case PA_WSTRING:                        /* All pointers are the same */
        T (PA_POINTER, pa_pointer, void *);
#undef T
